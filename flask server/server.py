from flask import Flask, request, jsonify
import os

app = Flask(__name__)

# Directory to save the received sensitive data
SAVE_FOLDER = 'received_data'
os.makedirs(SAVE_FOLDER, exist_ok=True)

@app.route('/test.html', methods=['POST'])
def receive_data():
    # Get the data from the request
    sensitive_data = request.form.get('data', '')

    if not sensitive_data:
        return jsonify({'error': 'No data received'}), 400

    # Save the received data to a file
    file_path = os.path.join(SAVE_FOLDER, 'sensitive_data.txt')
    with open(file_path, 'w') as file:
        file.write(sensitive_data)

    return jsonify({'message': 'Data successfully received'}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8000)
