from config_loader import Config
from data_processor import DataProcessor
import sys

def main():
    config_path = '.env'
    cfg = Config()
    try:
        cfg.load_config(config_path)
    except Exception as e:
        print(f"Failed to load config: {e}", file=sys.stderr)
        sys.exit(1)
    
    data_processor = DataProcessor(cfg)
    data_processor.process_data()

if __name__ == '__main__':
    main()
