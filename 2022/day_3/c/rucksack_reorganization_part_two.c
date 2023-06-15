#include <stdio.h>
#include <stdlib.h>


#define ITEMS_SIZE 52

typedef struct item_counter {
    short item;
    short count;
    short priority;
} ItemCounter; 


typedef struct item_counter_list {
    ItemCounter items[ITEMS_SIZE];
} ItemCounterList; 


// Initialize a rucksack counting list
void init_list(ItemCounterList * list) {
    // Initialize the list with items with lower priority (1-26)
    for (short i = 0; i < ITEMS_SIZE / 2; i++)
    {
        ItemCounter ic;
        ic.item = 97 + i;
        ic.priority = i + 1;  
        ic.count = 0; 
        list->items[i] = ic;
    }

    // Continue filling list with items with higher priority (27-52)
    short j = 0;
    for (short i = ITEMS_SIZE / 2; i < ITEMS_SIZE; i++)
    {
        ItemCounter ic;
        ic.item = 65 + j++;
        ic.priority = i + 1;  
        ic.count = 0; 
        list->items[i] = ic;
    }
}


// Find the badge that corresponds the counting list
short get_badge_prio(ItemCounterList * list) {
    for (int i = 0; i < ITEMS_SIZE; i++)
    {
        if (list->items[i].count == 3) {
            return list->items[i].priority; 
        }   
    }
}


// Count items in a line based on its index. For example if it is the first 
// line of a group n_line should be 0 if it is the second line of a group n_line
// should be 1 etc... 
void count_for_group_line(ItemCounterList * list, char * line, short n_line) { 

    for (char * c = line; *c != '\0'; c++)
    {
        short item = *c;

        for (int i = 0; i < ITEMS_SIZE; i++)
        {
            if (list->items[i].item == item && list->items[i].count == n_line) {
                list->items[i].count++;
                break;
            }   
        }
    }
}

// Reset list counters
void reset(ItemCounterList * list) { 
    for (int i = 0; i < ITEMS_SIZE; i++)
    {
        list->items[i].count = 0;
    }
}


// Print list state for debugging
void print_list(ItemCounterList * list)  {
    for (int i = 0; i < ITEMS_SIZE; i++)
    {
        ItemCounter obj = list->items[i];
        printf("\nItem: %c\nPriority: %d\nCount: %d\n", obj.item, obj.priority, obj.count);
    }
}


int main(int argc, char **argv)
{
    int sum = 0;
    ItemCounterList * i_list = (ItemCounterList *) malloc (sizeof(ItemCounterList));
    init_list(i_list);

    // Open the file for reading
    FILE *fd = fopen(argv[1], "r");
    if (fd == NULL) {
        perror("Error opening file");
        return -1;
    }

    char * line = NULL;
    size_t len = 0;
    size_t line_size;

    short reset_counter = 0;
    puts("\n========== Start parsing racksuck ==========\n");
    while ((line_size = getline(&line, &len, fd)) != -1) {
        printf("Retrieved line of length %zu:\n", line_size);
        printf("%s\n", line);

        short half_size = 0; 
        if (line_size % 2 != 0) 
            half_size = (line_size-1)/2; 
        else
            half_size = line_size/2; 


        reset_counter++;
        switch (reset_counter)
        {
        case 1:
            count_for_group_line(i_list, line, 0);
            break;
        case 2:
            count_for_group_line(i_list, line, 1);
            break;
        case 3:
            count_for_group_line(i_list, line, 2);
            print_list(i_list);
            sum += get_badge_prio(i_list);
            reset(i_list);
            reset_counter = 0;
            break;
        }

        printf("\n==================\n\n");
    }

    fclose(fd); 
    free(i_list);

    printf("Priority sum: %d", sum);


    return 0;
}