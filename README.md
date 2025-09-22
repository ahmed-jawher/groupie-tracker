# 🎸 Groupie Tracker

A modern web application that displays and visualizes information about music bands and artists, featuring a sleek dark theme and smooth animations.

## 🌟 Features

- **Artist Information**
  - Artist profiles with images
  - Band members listing
  - Concert locations
  - First album dates
  - Concert dates
  - Related information

- **Error Handling**
  - Custom error pages with consistent styling
  - User-friendly error messages
  - Smooth navigation back to safety

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher
- A modern web browser

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/groupie-tracker.git
cd groupie-tracker
```

2. Run the application:
```bash
go run .
```

3. Open your browser and visit:
```
http://localhost:8080
```

## 🛠️ Technical Details

### Project Structure
```
groupie-tracker/
├── main.go         # Entry point and server setup
├── handlers.go     # HTTP request handlers
├── models.go       # Data models and API interactions
└── templates/      # HTML templates
    ├── index.html  # Home page
    ├── artist.html # Artist details page
    └── error.html  # Error page template
```

### Technologies Used

- **Backend**
  - Go (Golang)
  - Standard library (net/http)
  - HTML templates

- **Frontend**
  - HTML5
  - CSS3 (Modern features)
  - Responsive Design

## 🎨 Design Features

- Animated cards with hover effects
- Dynamic content loading
- Responsive grid layout
- Custom error pages
- Smooth transitions
- Glass-morphism effects
- Neon accent colors

## 🌐 API Integration

The application integrates with the provided API endpoints:
- Artists information
- Locations
- Dates
- Relations

## 📱 Responsive Design

The application is fully responsive and works on:
- Desktop computers
- Tablets
- Mobile phones
- Various screen sizes and orientations

## ⚡ Performance

- Optimized image loading
- Smooth animations
- Efficient error handling
- Fast page loads

## 🔒 Error Handling

Custom error pages for:
- 404 Not Found
- 400 Bad Request
- 500 Internal Server Error

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!
