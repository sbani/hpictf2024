import random
import time

class DataProcessor:
    def __init__(self, config):
        self.config = config
    
    def process_data(self):
        processing_interval = self.config.get('PROCESSING_INTERVAL')
        if not processing_interval:
            raise ValueError("Processing interval not set in configuration.")
        
        interval = float(processing_interval)
        log_file_path = self.config.get('LOG_FILE_PATH')
        if not log_file_path:
            raise ValueError("Log file path not set in configuration.")
        
        with open(log_file_path, 'a') as log_file:
            while True:
                data_item = self.generate_data()
                log_message = f"Processed data: {data_item} at {time.ctime()}\n"
                print(log_message, end='')
                log_file.write(log_message)
                time.sleep(interval)
    
    def generate_data(self):
        random_data = random.randint(1, 100)
        return f"data_value_{random_data}"
