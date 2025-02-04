# Real Image Challenge 2016

In the cinema business, a feature film is usually provided to a regional distributor based on a contract for exhibition in a particular geographical territory.

Each authorization is specified by a combination of included and excluded regions. For example, a distributor might be authorzied in the following manner:
```
Permissions for DISTRIBUTOR1
INCLUDE: INDIA
INCLUDE: UNITEDSTATES
EXCLUDE: KARNATAKA-INDIA
EXCLUDE: CHENNAI-TAMILNADU-INDIA
```
This allows `DISTRIBUTOR1` to distribute in any city inside the United States and India, *except* cities in the state of Karnataka (in India) and the city of Chennai (in Tamil Nadu, India).

At this point, asking your program if `DISTRIBUTOR1` has permission to distribute in `CHICAGO-ILLINOIS-UNITEDSTATES` should get `YES` as the answer, and asking if distribution can happen in `CHENNAI-TAMILNADU-INDIA` should of course be `NO`. Asking if distribution is possible in `BANGALORE-KARNATAKA-INDIA` should also be `NO`, because the whole state of Karnataka has been excluded.

Sometimes, a distributor might split the work of distribution amount smaller sub-distiributors inside their authorized geographies. For instance, `DISTRIBUTOR1` might assign the following permissions to `DISTRIBUTOR2`:

```
Permissions for DISTRIBUTOR2 < DISTRIBUTOR1
INCLUDE: INDIA
EXCLUDE: TAMILNADU-INDIA
```
Now, `DISTRIBUTOR2` can distribute the movie anywhere in `INDIA`, except inside `TAMILNADU-INDIA` and `KARNATAKA-INDIA` - `DISTRIBUTOR2`'s permissions are always a subset of `DISTRIBUTOR1`'s permissions. It's impossible/invalid for `DISTRIBUTOR2` to have `INCLUDE: CHINA`, for example, because `DISTRIBUTOR1` isn't authorized to do that in the first place. 

If `DISTRIBUTOR2` authorizes `DISTRIBUTOR3` to handle just the city of Hubli, Karnataka, India, for example:
```
Permissions for DISTRIBUTOR3 < DISTRIBUTOR2 < DISTRIBUTOR1
INCLUDE: HUBLI-KARNATAKA-INDIA
```
Again, `DISTRIBUTOR2` cannot authorize `DISTRIBUTOR3` with a region that they themselves do not have access to. 

We've provided a CSV with the list of all countries, states and cities in the world that we know of - please use the data mentioned there for this program. *The codes you see there may be different from what you see here, so please always use the codes in the CSV*. This Readme is only an example. 

Write a program in any language you want (If you're here from Gophercon, use Go :D) that does this. Feel free to make your own input and output format / command line tool / GUI / Webservice / whatever you want. Feel free to hold the dataset in whatever structure you want, but try not to use external databases - as far as possible stick to your langauage without bringing in MySQL/Postgres/MongoDB/Redis/Etc.

To submit a solution, fork this repo and send a Pull Request on Github. 

For any questions or clarifications, raise an issue on this repo and we'll answer your questions as fast as we can.


# Distributor Permissions System

## Overview
This project implements a **Distributor Permissions System** that determines if a distributor has permission to distribute in a given region. It processes region mappings, distributor permissions, and queries efficiently using structured data handling.

## Folder Structure
```
project-root/
│── main.go
│── data/
│   ├── cities.csv         # Region mappings (City, State, Country)
│   ├── permissions.txt    # Distributor permissions
│   ├── queries.txt        # Queries for permission checks
│── pkg/
│   ├── permissions/       # Handles distributor permissions
│   │   ├── permissions.go
│   ├── regions/           # Handles region mappings
│   │   ├── regions.go
│   ├── queries/           # Handles query processing
│   │   ├── queries.go
```

## Installation & Usage
### Prerequisites
- Go 1.18 or later installed

### Steps to Run
1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/distributor-permissions.git
   cd distributor-permissions
   ```
2. Place the required CSV and text files inside the `data/` directory.
3. Run the program:
   ```sh
   go run main.go
   ```

## Data Files
### `cities.csv`
Contains location mappings in the format:
```
ID, Name, Type, City, State, Country
```
Example:
```
1, New York, City, New York, NY, USA
```

### `permissions.txt`
Defines permissions for distributors.
```
Permissions for Distributor1
INCLUDE: New York-NY-USA
EXCLUDE: NY-USA
Permissions for Distributor2 < Distributor1
```

### `queries.txt`
Contains queries in the format:
```
Can Distributor1 distribute in New York-NY-USA?
```

## Features
- **Hierarchical Permissions**: Distributors can inherit permissions from parent distributors.
- **Region Mapping**: Matches city, state, and country names efficiently.
- **Query Processing**: Checks whether a distributor can operate in a given region.

## License
This project is licensed under the MIT License.

## Author
[Your Name]




