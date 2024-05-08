# Gossiper

Gossiper is a simple implementation of a gossip protocol in Go. It demonstrates how nodes in a distributed system can exchange information and reach a consistent state through a peer-to-peer communication model.

## Overview

In a gossip protocol, each node periodically communicates with a randomly selected neighbor, exchanging state information. Through this process, information spreads throughout the network, eventually leading to a consistent state across all nodes.

This project consists of the following packages:

- `types`: Defines the basic types used in the gossip protocol, such as `Node` and `State`.
- `network`: Provides functionality for selecting a random neighbor and determining the gossip interval.
- `node`: Contains the core logic for the gossip protocol, implemented in the `Gossip` function.

## Usage

To run the gossip protocol simulation, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gossiper.git
   ```
2. Navigate to the project directory:
   ```bash
   cd gossiper
   ```
3. Run the main program:
   ```bash
   go run main.go
   ```
The program will create a set of nodes, set up their neighbor connections, and start the gossip protocol. It will periodically print the state of each node, allowing you to observe how the state changes propagate through the network.

## Customization

You can customize the behavior of the gossip protocol by modifying the following:

- `types.Node`: Adjust the structure of the `Node` type to include additional properties or methods.
- `types.State`: Modify the `State` type to represent the specific state information exchanged between nodes.
- `network.GossipInterval`: Change the function to determine the interval between gossip rounds.
- `network.SelectRandomNeighbor`: Implement a different strategy for selecting a random neighbor.
- `node.Gossip`: Modify the gossip logic to incorporate additional behavior or conditions.

Feel free to experiment and extend the project to suit your specific requirements.

## Contributing

Contributions to Gossiper are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/devlongs/gossiper).

