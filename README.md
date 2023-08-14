# Kubetune

Kubetune is a command-line tool for managing Kubernetes configuration files (kubeconfig). It allows users to view, edit, and track changes to the kubeconfig file in a user-friendly way.

## Features

- **View**: Display the current kubeconfig file.
- **Change**: Open the kubeconfig file in a text editor to make changes.
- **Diff**: Compare the current kubeconfig file with a previous version and display the differences.

## Installation

To install Kubetune, you can clone the repository and build the tool using Go:

git clone https://github.com/zacharyrgonzales-portfolio/kubetune
cd kubetune
go build

## Usage

Here's how to use the main commands of Kubetune:

### View the kubeconfig file

kubetune view

### Edit the kubeconfig file

kubetune change

### Show differences between the current and previous versions of the kubeconfig file

kubetune diff

## Contributing

If you'd like to contribute to Kubetune, please open an issue or submit a pull request.

## License

Kubetune is released under the MIT License. See the [LICENSE](LICENSE) file for details.
