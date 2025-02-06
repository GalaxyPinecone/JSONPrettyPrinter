import json

json_data = input("Enter JSON string: ")

try:
    parsed = json.loads(json_data)
    print(json.dumps(parsed, indent=4))
except json.JSONDecodeError:
    print("Invalid JSON data.")
