#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
    if (argc < 2) {
        printf("Provide the input path as a parameter\n");
        return 0;
    }

    // Open the file for reading
    FILE *fd = fopen(argv[1], "r");
    if (fd == NULL) {
        perror("Error opening file");
        return -1;
    }

    // Read the file line by line
    // TODO: For part two, fix bug in the case that one of the top three is after the max in will not be registered 
    char line[100];
    int max_calories = 0;
    int new_max_count = 0;
    int current_calories = 0;
    int max_three_calories[] = {0, 0, 0};
    while (fgets(line, 100, fd) != NULL) {
        current_calories += atoi(line);
        if (line[0] == '\n') {
            if (max_calories < current_calories) {
                max_calories = current_calories;
                max_three_calories[new_max_count++ % 3] = max_calories; 
                printf("Found an elf with more calories: %d\n", max_calories);
            }
            current_calories = 0;
        }
    }

    int max_three_sum = 0;
    for (size_t i = 0; i < sizeof(max_three_calories) / sizeof(int); i++)
    {
        max_three_sum += max_three_calories[i];
    }
    
    printf("The elf with the most calories has: %d\n", max_calories);
    printf("The three elves with the most calories have in total: %d\n", max_three_sum);

    // Close the file
    fclose(fd);

    return 0;
}