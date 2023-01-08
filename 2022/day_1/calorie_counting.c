#include <stdio.h>
#include <stdlib.h>

// compare values with descending order
int cmpfunc (const void * a, const void * b) {
   return ( *(int*)b - *(int*)a);
}

int main(int argc, char **argv)
{
    // Open the file for reading
    FILE *fd = fopen("input.txt", "r");
    if (fd == NULL) {
        perror("Error opening file");
        return -1;
    }

    // Read the file line by line
    char line[100];
    int calorie_sum_count = 0;
    int current_calories = 0;
    int calorie_sums[1024] = {0};
    int max_three_calories[3] = {0};

    while (fgets(line, 100, fd) != NULL) {
        current_calories += atoi(line);
        if (line[0] == '\n') {
            calorie_sums[calorie_sum_count++] = current_calories;
            current_calories = 0;
        }
    }

    qsort(calorie_sums, sizeof(calorie_sums) / sizeof(int), sizeof(int), cmpfunc);

    int max_three_sum = 0;
    for (size_t i = 0; i < 3; i++)
    {
        max_three_sum += calorie_sums[i];
    }
    
    printf("The elf with the most calories has: %d\n", calorie_sums[0]);
    printf("The three elves with the most calories have in total: %d\n", max_three_sum);

    // Close the file
    fclose(fd);

    return 0;
}