# 🌐 **Aladia Project – Test Submission**  

This repository presents a complete implementation of key user authentication flows and landing experiences, built using **Nuxt.js 3**. Below is a concise explanation of the technologies and tools used, along with the design decisions made throughout the project. 

---

## ✅ **Implemented Features**  
- **Signup/Login**: Supports both registration and login with email and password, handling all essential states, including errors (e.g., invalid credentials, existing accounts).  
- **Third-Party Login (Simulated)**: Provides visual flows for Google, Facebook, and Apple login states, ensuring UI readiness without real integration.  
- **Password Recovery and Reset**: Separate flows for resetting passwords, including cases for Google, Facebook, and Apple users.  
- **Logout**: Ensures proper session cleanup and state management.  
- **Landing Page**: A responsive landing page to guide users into the authentication flows.

---

## 📘 **Storybook and Chromatic Integration**  
We used **Storybook** to develop, document, and test components in isolation. This allowed us to visualize each component state, such as loading, success, and error, ensuring they behave correctly. **Chromatic** was integrated for visual regression testing to detect unintended UI changes with every update, keeping the design consistent throughout development.  

---

## 🛡️ **Cypress for End-to-End Testing**  
**Cypress** was used to automate E2E tests across key user flows, covering signup, login, third-party login simulations, password recovery, and logout. This allowed us to ensure that the application performs reliably under different conditions. **Pointer events** (`@pointerenter`, `@pointerup`) were used instead of traditional click or touch events, unifying interactions across all devices (mouse, touch, or stylus). This decision streamlined our testing process and improved cross-device compatibility.

---

## 🔄 **CI/CD Pipeline with GitHub Actions**  
A **CI/CD pipeline** was implemented using **GitHub Actions**. This pipeline automatically triggers Cypress and Chromatic tests for every code change, ensuring that both functionality and design remain intact. This setup allows us to catch issues early in the development cycle and maintain consistent code quality throughout.  

---

## 📂 **State Management with store.js**  
**store.js** was used for state management to ensure smooth data flow between components and across pages. This was crucial for maintaining session states during login, signup, and logout processes, providing a seamless user experience without unexpected behavior or data loss.

---

## ⚠️ **Error Handling Across Flows**  
Every flow includes thorough error handling to cover edge cases. For example, the login and signup processes respond to invalid credentials or existing accounts with clear feedback. These error states are documented in Storybook, ensuring they are tested and visually correct. 

---

## 🎯 **Summary**  
The Aladia project combines modern tools and thoughtful design to build a robust and maintainable application. With **Storybook** and **Chromatic** driving component-level testing, **Cypress** ensuring real-world reliability, and a well-organized CI/CD pipeline, the project is built to be scalable and easy to extend. By leveraging **store.js** for state management and focusing on error handling, we created a seamless and predictable user experience across all key flows.
