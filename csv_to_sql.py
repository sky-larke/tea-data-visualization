import csv
import re


# Read the CSV file and convert to SQL insert statements
def convert_csv_to_sql(csv_file, output_sql_file):
    with open(csv_file, mode='r', newline='', encoding='utf-8') as infile, open(output_sql_file, mode='w', encoding='utf-8') as outfile:
        reader = csv.DictReader(infile, delimiter=',', quotechar='"')  # Handles commas and quotes correctly
        
        # Strip spaces from the field names
        fieldnames = [field.strip() for field in reader.fieldnames]
        
        # Loop through each row in the CSV file
        rank = 0
        for row in reader:
            # Strip spaces from the values
            name = row['Name'].strip()
            vendor = row['Vendor'].strip()
            year = row['Year'].strip()
            category = row['Category'].strip()
            subcategory = row['Sub-category'].strip() if row['Sub-category'].strip() else 'NULL'
            cost = row['Cost'].strip()
            amount = row['Amount'].strip()
            rank += 1
            
            # Extract vendor from the name (if available)
            
            # Format the SQL insert statement
            sql = f"INSERT INTO teas(rank, vendor, name, year, type, subtype, cost, amount) VALUES ('{rank}', '{vendor}', '{name}', '{year if year else 0}', '{category}', '{subcategory}', '{cost}', '{amount}');\n"
            
            # Write to the output SQL file
            outfile.write(sql)

# Example usage
csv_file = 'tea.csv'  # Path to the input CSV file
output_sql_file = 'teas_insert.sql'  # Path to the output SQL file

convert_csv_to_sql(csv_file, output_sql_file)
