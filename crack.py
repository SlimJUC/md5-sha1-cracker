import hashlib
import concurrent.futures

# Generate hash for a password
def generate_hash(password, algorithm):
    if algorithm == 'sha1':
        h = hashlib.sha1()
    elif algorithm == 'md5':
        h = hashlib.md5()
    else:
        raise ValueError(f"Unsupported algorithm: {algorithm}")
    h.update(password.encode('utf-8'))
    return h.hexdigest()

# Get user input for stored password hash and read user passwords from a file
stored_password_hash = input("Enter the stored password hash: ")
if len(stored_password_hash) == 32:
    algorithm = 'md5'
elif len(stored_password_hash) == 40:
    algorithm = 'sha1'
else:
    raise ValueError("Unable to detect hash algorithm")

password_file = input("Enter the file name containing the list of passwords: ")

# Read the passwords from the file
with open(password_file, 'r', encoding='iso-8859-1') as f:
    passwords = [line.strip() for line in f.readlines()]

# Check each password and output the result
with concurrent.futures.ThreadPoolExecutor() as executor:
    results = []
    for password in passwords:
        results.append(executor.submit(generate_hash, password, algorithm))

    for i, result in enumerate(concurrent.futures.as_completed(results)):
        password_hash = result.result()
        if password_hash == stored_password_hash:
            print(f"Password {passwords[i]} is correct")
            break
    else:
        print("None of the passwords in the list are correct")
