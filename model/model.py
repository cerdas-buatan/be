import tensorflow as tf
from transformers import TFBertForSequenceClassification, BertTokenizer
import pandas as pd
import os

# Absolute path to the dataset
dataset_path = os.path.abspath('dataset-a.csv')

# Attempt to load the dataset with the correct delimiter
try:
    df = pd.read_csv(dataset_path, delimiter='|')
except pd.errors.ParserError as e:
    print(f"Error reading CSV file: {e}")
    print("Please check your CSV file for formatting issues.")
    raise

# Print the column names and preview the first few rows to verify
print("Columns in the dataset:", df.columns)
print(df.head())

# Ensure that 'question' and 'answer' columns exist
expected_columns = ['question', 'answer']
missing_columns = [col for col in expected_columns if col not in df.columns]
if missing_columns:
    print(f"The dataset is missing the following columns: {', '.join(missing_columns)}")
    raise KeyError("Required columns are missing in the dataset.")

# Prepare the dataset
tokenizer = BertTokenizer.from_pretrained('indobenchmark/indobert-base-p2')
input_ids = []
attention_masks = []

for question in df['question']:
    encoded = tokenizer.encode_plus(
        question,
        add_special_tokens=True,
        max_length=64,
        padding='max_length',
        truncation=True,
        return_attention_mask=True
    )
    input_ids.append(encoded['input_ids'])
    attention_masks.append(encoded['attention_mask'])

input_ids = tf.constant(input_ids)
attention_masks = tf.constant(attention_masks)
labels = tf.constant(df['answer'].values)

# Create TensorFlow Dataset
dataset = tf.data.Dataset.from_tensor_slices(({"input_ids": input_ids, "attention_mask": attention_masks}, labels))
dataset = dataset.shuffle(len(df)).batch(32)

# Load model
model = TFBertForSequenceClassification.from_pretrained('indobenchmark/indobert-base-p2')

# Compile the model
model.compile(
    optimizer=tf.keras.optimizers.Adam(learning_rate=5e-5),
    loss=tf.keras.losses.SparseCategoricalCrossentropy(from_logits=True),
    metrics=['accuracy']
)

# Train the model
model.fit(dataset, epochs=3)

# Save the model
model.save_pretrained('indobert_model')
