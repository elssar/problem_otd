"""
Solution for Problem of the Day for 2014-05-13, File Stats
http://www.problemotd.com/problem/file-stats/
"""

import os
import sys
import time


def get_file_stats():
    try:
        path = sys.argv[1]
        with open(path, 'r') as file_obj:
            words = file_obj.read()
        os_stats = os.stat(path)
        return words, os_stats
    except IndexError:
        print "Error: Need a file path"
        sys.exit(1)
    except IOError, err:
        print "Error: {0}".format(err)
        sys.exit(1)
    except Exception, err:
        print "Error: {0}".format(err)


def print_results(words, stats):
    num_words = len(words.split(" "))
    num_lines = words.count('\n')
    filename = os.path.split(sys.argv[1])[-1]
    print "\nStats for \"{}\":\n".format(filename)
    print ("Words: {}, Lines: {}, Size: {} bytes,\n"
          "Created: {},\nLast Modified: {}".format(
                num_words,
                num_lines,
                stats.st_size,
                time.ctime(stats.st_ctime),
                time.ctime(stats.st_mtime)))


def main():
    words, stats = get_file_stats()
    print_results(words, stats)


if __name__ == "__main__":
    main()
