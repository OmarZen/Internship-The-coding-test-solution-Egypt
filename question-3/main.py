import psycopg2
import json

# Establish a connection to the PostgreSQL database
conn = psycopg2.connect(
    host="localhost",
    port="5432",
    database="your_database",
    user="your_username",
    password="your_password"
)

# Create a cursor object to interact with the database
cur = conn.cursor()

# Execute the query to fetch data from the user_table
cur.execute("SELECT * FROM public.user_table")
rows = cur.fetchall()

users = []

# Iterate over the query result and populate the users list
for row in rows:
    user = {
        "user_id": row[0],
        "name": row[1],
        "age": row[2],
        "phone": row[3] if row[3] else None
    }
    users.append(user)

# Create a dictionary with the fetched data
response = {
    "status_code": 200,
    "data": users
}

# Convert the response dictionary to JSON
jsonData = json.dumps(response)

# Print the JSON data
print(jsonData)

# Close the cursor and connection
cur.close()
conn.close()
