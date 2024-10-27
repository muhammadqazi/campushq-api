# Contributing to CampusHQ API

Thank you for your interest in contributing to CampusHQ API! We aim to make campus management easier and more efficient through this API.

## Project Overview

CampusHQ API is a backend service designed to support campus management operations. The project follows clean architecture principles to maintain code quality and scalability.

## Getting Started

1. **Fork and Clone**
   ```bash
   git clone https://github.com/muhammadqazi/campushq-api
   cd campushq-api
   ```

2. **Install Dependencies**
   ```bash
   go mod download
   ```

3. **Set Up Environment**
   - Copy `.env.example` to `.env`
   - Update environment variables as needed

## Development Guidelines

1. **Code Style**
   - Follow standard Go conventions
   - Use meaningful variable and function names
   - Comment complex logic

2. **Git Workflow**
   - Create feature branches from `master`
   - Branch naming: `feature/feature-name` or `fix/bug-name`
   - Keep commits focused and meaningful

3. **Commit Messages**
   ```
   type(scope): brief description

   Detailed description if needed
   ```
   Types: feat, fix, docs, style, refactor, test, chore

## Pull Request Process

1. **Before Submitting**
   - Ensure all tests pass
   - Update documentation if needed
   - Test your changes thoroughly
   - Pull latest changes from master

2. **PR Requirements**
   - Clear description of changes
   - Link to related issues
   - Screenshots for UI changes (if applicable)
   - Updated tests

3. **Review Process**
   - CI checks must pass
   - One approved review required
   - No merge conflicts

## Testing

1. **Running Tests**
   ```bash
   go test ./...
   ```

2. **Test Coverage**
   ```bash
   go test -cover ./...
   ```


## Need Help?

- Check existing issues first
- For quick questions, use discussions
- Tag maintainers when needed

## License

By contributing, you agree that your contributions will be licensed under the project's license.

## Code of Conduct

- Be respectful and inclusive
- Focus on constructive feedback
- Help others learn and grow

Thank you for contributing to making campus management better! 🎓
