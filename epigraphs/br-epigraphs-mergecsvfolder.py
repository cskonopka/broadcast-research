import os
import pandas as pd
import re
from datetime import datetime

def get_month_from_filename(filename):
    # Extract month from filename using regex (e.g., "data_january.csv" or "2023-01_sales.csv")
    month_patterns = [
        r'january|jan', r'february|feb', r'march|mar', r'april|apr', 
        r'may', r'june|jun', r'july|jul', r'august|aug', 
        r'september|sep', r'october|oct', r'november|nov', r'december|dec'
    ]
    filename_lower = filename.lower()
    for i, pattern in enumerate(month_patterns):
        if re.search(pattern, filename_lower):
            return (month_patterns[i].split('|')[0]).capitalize()
    # Try to extract from date-like patterns (e.g., "2023-01")
    date_pattern = r'\d{4}[-_](0[1-9]|1[0-2])'
    match = re.search(date_pattern, filename_lower)
    if match:
        month_num = int(match.group(1))
        return datetime.strptime(str(month_num), '%m').strftime('%B')
    return "Unknown"

def split_camel_case(phrase):
    # Split camelCase or PascalCase into words, preserving original case
    # e.g., "fuzzyOvals" -> "fuzzy Ovals", "pinkBars" -> "pink Bars"
    if not phrase:
        return phrase
    # Use regex to insert space before each capital letter or number
    # This captures the position before a capital letter or number, ensuring no split on lowercase
    split_phrase = re.sub(r'(?<!^)(?=[A-Z0-9])', ' ', phrase)
    return split_phrase

def merge_csv_files(input_dir, output_file):
    # Check if directory exists
    if not os.path.isdir(input_dir):
        print(f"Error: Directory {input_dir} does not exist.")
        return
    
    # Get list of CSV files
    csv_files = [f for f in os.listdir(input_dir) if f.endswith('.csv')]
    if not csv_files:
        print(f"No CSV files found in {input_dir}")
        return
    
    # List to store individual dataframes
    dfs = []
    
    # Process each CSV file
    for file in csv_files:
        file_path = os.path.join(input_dir, file)
        try:
            # Read CSV file, assuming first two columns are Date and Name
            df = pd.read_csv(file_path, usecols=[0, 1], names=['Date', 'Name'], header=None)
            # Drop rows where Date or Name is missing
            df = df.dropna(subset=['Date', 'Name'])
            # Add Month column based on filename
            month = get_month_from_filename(file)
            df['Month'] = month
            dfs.append(df)
            print(f"Processed {file} - Month: {month}, Rows: {len(df)}")
        except Exception as e:
            print(f"Error processing {file}: {str(e)}")
    
    if not dfs:
        print("No valid CSV files were processed.")
        return
    
    # Merge all dataframes
    merged_df = pd.concat(dfs, ignore_index=True)
    
    # Ensure only the desired columns are included
    merged_df = merged_df[['Date', 'Name', 'Month']]
    
    # Convert Date column to datetime for proper sorting
    try:
        merged_df['Date'] = pd.to_datetime(merged_df['Date'], format='%m-%d-%Y')
    except Exception as e:
        print(f"Error converting dates: {str(e)}")
        return
    
    # Sort by Date
    merged_df = merged_df.sort_values(by='Date')
    
    # Convert Date back to original string format for output
    merged_df['Date'] = merged_df['Date'].dt.strftime('%m-%d-%Y')
    
    # Add ProcessedName column as a copy of Name
    merged_df['ProcessedName'] = merged_df['Name']
    
    # Apply camelCase/PascalCase splitting to ProcessedName, preserving original case
    merged_df['ProcessedName'] = merged_df['ProcessedName'].apply(split_camel_case)
    
    # Save to output file
    try:
        merged_df.to_csv(output_file, index=False)
        print(f"Successfully merged {len(dfs)} files into {output_file}")
        print(f"Total rows: {len(merged_df)}")
    except Exception as e:
        print(f"Error saving merged file: {str(e)}")

if __name__ == "__main__":
    # Example usage
    input_directory = "/Users/rr/Documents/SPLOOSH/broadcast-research/epigraphs/2019"  # Replace with your directory path
    output_file = "merged_output.csv"      # Output file name
    merge_csv_files(input_directory, output_file)