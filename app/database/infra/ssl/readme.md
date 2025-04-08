# Database SSL Credentials

### Example and Test Setup 🎯
The contents in the `./database/ssl` folder are provided as an example and for testing
purposes. They demonstrate how SSL (Secure Sockets Layer) can be used to establish a
secure connection with a PostgreSQL database. This setup is meant to give you a clear
idea of how to implement SSL, without affecting your production environment.

### Why Use SSL? 🔒
SSL is a critical security measure that encrypts the connection between your application
and the database. Here’s why it’s important:
- **Data Protection:** SSL helps protect sensitive information like credentials and database queries from unauthorized access.
- **Secure Communication:** Encrypting the data ensures that even if someone intercepts the traffic, they won’t be able to read it.
- **Trust and Reliability:** Using SSL shows that you take security seriously, which can build trust with your users and stakeholders.

### Creating Your Own Secure Credentials 🔧
To create your own secure database credentials:
1. Open your terminal.
2. Run the below command in the root folder of your project.
3. Follow the on-screen instructions to generate your unique credentials.
```bash
  make db:secure
```

### Summary 🎉
- **Example Setup:** The ./database/ssl folder is only for demonstration and testing.
- **SSL Benefits:** Encrypts your connection to protect sensitive data.
- **Setup Process:** Use make db:secure to generate your own secure credentials.

Embrace these practices to keep your data safe and secure.
Happy coding! 😃
