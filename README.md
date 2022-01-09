# Online Store
## Problem
> 1 Our inventory quantities are often misreported, and some items even go as far as having a negative inventory quantity. 
2 The misreported items are those that performed very well on our 12.12 event. 3 Because of these misreported inventory quantities, the Order Processing department was unable to fulfill a lot of orders, and thus requested help from our Customer Service department to call our customers and notify them that we have had to cancel their orders.

The problem is caused by multiple users accessing the data and performing update/write operations to the system at the same time. What is commonly known as a race condition.
The API application that has been made validates whether the stock is still there or not when there is a request to create or update.

## ERD
![image](https://user-images.githubusercontent.com/23447522/148683649-3c8c0198-35c9-4f72-ac7b-3aae2d6d8b61.png)

# Treasure Hunt
Please find in Maze folder

![image](https://user-images.githubusercontent.com/23447522/148683613-954b0cb7-35e0-434b-af11-b78bc5a6ba8e.png)
