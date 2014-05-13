"""
Solution for Problem of the Day for 2014-05-12, Character Frequency
http://www.problemotd.com/problem/character-frequency/
"""

from collections import Counter
import operator


def get_sorted_freq(text):
    freq = Counter(text)
    return sorted(freq.iteritems(), key=operator.itemgetter(1), reverse=True)


def print_result(sorted_freq):
    for char, freq in sorted_freq:
        print "{0}\t{1}".format(char, "." * freq)


def main():
    text = raw_input("Enter text: ")
    sorted_freq = get_sorted_freq(text)
    print_result(sorted_freq)


if __name__ == "__main__":
    main()
