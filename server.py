# server.py
from flask import Flask, request, jsonify
from transformers import T5ForConditionalGeneration, T5Tokenizer


app = Flask(__name__)
location = "./t5_small_model/t5_small_model_final"
model = T5ForConditionalGeneration.from_pretrained(location)
tokenizer = T5Tokenizer.from_pretrained(location)


@app.route('/query', methods=['POST'])
def query():
    data = request.get_json(force=True)
    queryText = data.get("query", "")
    if not queryText.lower().startswith("translate english to bash:"):
        queryText = "translate English to Bash: " + queryText

    inputs = tokenizer.encode(queryText, return_tensors="pt", truncation=True, 
                                padding="max_length", max_length=64)
    outputs = model.generate(inputs, max_length=64, num_beams=5, early_stopping=True)
    command = tokenizer.decode(outputs[0], skip_special_tokens=True)
    print("Generated Command:", command)  # Debug 
    return jsonify({"command": command})

if __name__ == '__main__':
    app.run(host="0.0.0.0", port=5000)
