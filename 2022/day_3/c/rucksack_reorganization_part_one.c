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

void inc_item_comp_first(ItemCounterList * list, short item) { 
    for (int i = 0; i < ITEMS_SIZE; i++)
    {
        if (list->items[i].item == item && list->items[i].count == 0) {
            list->items[i].count++;
            break;
        }   
    }
}

void inc_item_comp_second(ItemCounterList * list, short item) { 
    for (int i = 0; i < ITEMS_SIZE; i++)
    {
        if (list->items[i].item == item && list->items[i].count == 1) {
            list->items[i].count++;
            break;
        }   
    }
}

void reset(ItemCounterList * list) { 
    for (int i = 0; i < ITEMS_SIZE; i++)
    {
        list->items[i].count = 0;
    }
}

int dub_prio_sum(ItemCounterList * list) { 
    int sum = 0;
    for (int i = 0; i < ITEMS_SIZE; i++)
    {
        if (list->items[i].count == 2) {
            sum += list->items[i].priority;
        }
    }
    return sum;
}

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

    puts("\n========== Start parsing racksuck ==========\n");
    while ((line_size = getline(&line, &len, fd)) != -1) {
        // printf("Retrieved line of length %zu:\n", line_size);
        // printf("%s\n", line);

        short half_size = 0; 
        if (line_size % 2 != 0) 
            half_size = (line_size-1)/2; 
        else
            half_size = line_size/2; 
        

        // printf("\nFirst compartment \n");
        for (int i = 0; i < half_size ; i++)
        {
            inc_item_comp_first(i_list, (short)line[i]);
            // printf("%c : ", line[i]);
        }

        // printf("\nSecond compartment \n");
        for (int i = half_size; i < line_size; i++)
        {
            inc_item_comp_second(i_list, (short)line[i]);
            // printf("%c : ", line[i]);
        }

        // printf("\n==================\n\n");

        // print_list(i_list);
        sum += dub_prio_sum(i_list);
        reset(i_list);
    }

    fclose(fd); 
    free(i_list);

    printf("Duplicate priority sum: %d", sum);


    return 0;
}