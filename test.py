from transformers import T5ForConditionalGeneration, T5Tokenizer

model_dir = "./t5_small_model/t5_small_model_final"
model = T5ForConditionalGeneration.from_pretrained(model_dir)
tokenizer = T5Tokenizer.from_pretrained(model_dir)

input_text = "translate English to Bash: find all jpg files modified last week"
print("Prompt:", input_text)
inputs = tokenizer.encode(input_text, return_tensors="pt", truncation=True, padding="max_length", max_length=64)
outputs = model.generate(inputs, max_length=64, num_beams=5, early_stopping=False)
print("Raw Generated IDs:", outputs)
generated_command = tokenizer.decode(outputs[0], skip_special_tokens=True)
print("Generated Command:", generated_command)
