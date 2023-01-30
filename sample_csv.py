import numpy as np
import csv

def sample_csv():
    sampling_freq , amp , bias = 44100, 1.5, 2.0
    t = np.arange(0, 0.5, 1/sampling_freq)
    sin1 = amp * np.sin(2 * np.pi * 400 * t) + bias
    sin2 = amp * np.sin(2 * np.pi * 1000 * t) + bias
    sin3 = amp * np.sin(2 * np.pi * 3000 * t) + bias
    combined = [a1 + a2 + a3 for (a1, a2, a3) in zip(sin1, sin2, sin3)]
    
    with open('sample.csv', 'w', newline="") as f:
        writer = csv.writer(f)
        for r2 in range(int(len(sin1))):
            writer.writerow([combined[r2]])

if __name__ == '__main__':
    sample_csv()