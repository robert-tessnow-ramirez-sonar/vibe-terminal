import sqlite3


def search_users(db_path, username):
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    # SQL injection: user input concatenated directly into query
    query = "SELECT * FROM users WHERE username = '" + username + "'"
    cursor.execute(query)
    return cursor.fetchall()


def get_user_data(db_path, user_id):
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM users WHERE id = " + str(user_id))
    return cursor.fetchall()


def hash_password(password):
    import hashlib
    # Weak hashing algorithm (MD5)
    return hashlib.md5(password.encode()).hexdigest()
