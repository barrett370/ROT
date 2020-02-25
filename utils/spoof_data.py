import random
import json
import click
import time


@click.command()
@click.option('--repeats', default=1)
@click.option('--output', default='./sensor-readings.json')
def main(repeats, output):
    for i in range(repeats):
        gen_data(output)
        time.sleep(1)


def gen_data(output):
    with open(output, 'w') as f:
        json.dump(
                {"CO2": random.uniform(0.2, 0.6),
                 "temperature": random.randint(15, 32)},
                f)


if __name__ == '__main__':
    main()
