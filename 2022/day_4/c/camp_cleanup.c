#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define bool int
#define true 1
#define false 0

typedef struct pair {
    int x;
    int y;
} Pair;

typedef struct sections {
    Pair first;
    Pair second;
} Sections;


void print_sections(Sections * sections){
    printf("Sections 1: %d-%d\nSections 2: %d-%d\n",
    sections->first.x, sections->first.y, sections->second.x, sections->second.y);
}

Sections * split_sections(char * line) {
    Sections * sections = (Sections *) malloc (sizeof(Sections));

    sections->first.x = atoi(strtok(line, ",-"));
    sections->first.y = atoi(strtok(NULL, ",-"));

    sections->second.x = atoi(strtok(NULL, ",-"));
    sections->second.y = atoi(strtok(NULL, ",-"));

    return sections; 
}

bool overlapping_sections(Sections * sections) {
    if (sections->first.x <= sections->second.x && sections->first.y >= sections->second.y) {
        return true;
    } 
    else if (sections->second.x <= sections->first.x && sections->second.y >= sections->first.y) {
        return true;
    } 
    else {
        return false;
    }
}

bool intersecting_sections(Sections * sections) {
    int count = 0;

    for (size_t i = sections->first.x; i <= sections->first.y; i++)
        for (size_t j = sections->second.x; j <= sections->second.y; j++)
            if (i == j) count++;
        
    return count > 0;
}

int main(int argc, char **argv)
{
    // Open the file for reading
    FILE *fd = fopen(argv[1], "r");
    if (fd == NULL) {
        perror("Error opening file");
        return -1;
    }

    char * line = NULL;
    size_t len = 0;
    size_t line_size;

    int overlapping_sum = 0;
    int intersecting_sum = 0;

    while ((line_size = getline(&line, &len, fd)) != -1) {
        Sections * sections = split_sections(line);
        print_sections(sections);
        bool overlapping = overlapping_sections(sections);
        overlapping_sum += overlapping;

        bool intersecting = intersecting_sections(sections);
        intersecting_sum += intersecting;
        printf("Overlapping section: %d\nIntersecting section: %d\n===================\n\n", overlapping, intersecting);
    }

    printf("There are %d overlapping sections\n", overlapping_sum);
    printf("There are %d intersecting sections\n\n", intersecting_sum);
    return 0;
}