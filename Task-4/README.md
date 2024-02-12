# Go Tasks

## Task-4 : JSON Marshaling and Unmarshaling

### Instructions

1. Given 3 JSON files: tech.json, contacts.json, user.json.
2. Unmarshal tech.json, contact.json and user.json and store that data into the struct.
3. Using those struct, Convert that data into the given final output stuct. And marshal it into the Given output JSON format.
4. Write that final JSON data into the final-output.json file.

### Solution

1. Read all given JSON files and unmarshal those data into the struct.
2. Then create one final output as per desired JSON format, and append it in the slice of FinalOutput struct.
3. Marshal that slice of FinalOutput struct and achieve one final JSON string.
4. Write that final JSON string into the final-output.json file using os module.

### How to Run Code ?

1. Clone this Repo
2. Go to `task-4` branch
3. `cd task-4` to switch directory
4. `go run main.go` to execute the file