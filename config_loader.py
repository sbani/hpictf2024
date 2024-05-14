import os

class Config:
    def __init__(self):
        self.data = {}

    def load_config(self, file_path):
        if not os.path.exists(file_path):
            raise FileNotFoundError(f"No such file: '{file_path}'")
        
        with open(file_path, 'r') as file:
            for line in file:
                line = line.strip()
                if line and not line.startswith('#'):
                    key, value = line.split('=', 1)
                    self.data[key.strip()] = value.strip()

    def get(self, key):
        return self.data.get(key, None)
