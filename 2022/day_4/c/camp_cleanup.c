#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct pair {
    int x;
    int y;
} Pair;

typedef struct sections {
    Pair first;
    Pair second;
} Sections;

// void print_sections(char ** sections){
//     printf("Sections 1: %s\nSections 2: %s\n===================\n\n", sections[0], sections[1]);
// }

void print_sections(Sections * sections){
    printf("Sections 1: %d-%d\nSections 2: %d-%d\n===================\n\n", sections->first.x, sections->first.y, sections->second.x, sections->second.y);
}

Sections * split_sections(char * line) {
    Sections * sections = (Sections *) malloc (sizeof(Sections));

    sections->first.x = atoi(strtok(line, ",-"));
    sections->first.y = atoi(strtok(NULL, ",-"));

    sections->second.x = atoi(strtok(NULL, ",-"));
    sections->second.y = atoi(strtok(NULL, ",-"));

    return sections; 
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

    while ((line_size = getline(&line, &len, fd)) != -1) {
        printf("Retrieved line of length %zu:\n", line_size);
        printf("%s\n", line);

        Sections * sections = split_sections(line);
        print_sections(sections);

    }


    return 0;
}