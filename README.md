🌐 # Aladia Website




 a modern web project built using the power of Nuxt.js for seamless performance and maintainability. With robust testing integrations such as Storybook and Cypress, we ensure that each component and feature works beautifully, backed by Chromatic's visual regression for perfect UI across changes.

🛠️ ## Tech Stack

Framework: Nuxt.js 3
Component Library: Storybook
Visual Regression Testing: Chromatic
End-to-End Testing: Cypress
🚀 ## Project Setup

1️⃣ ### Install Dependencies
Ensure you have Node.js installed, then run:

**npm install**
2️⃣ ### Run the Development Server
Start your local development server with:

**npm run dev**
This will launch the site at http://localhost:3000/.

3️⃣ ### Storybook Setup
We use Storybook to develop, document, and test UI components in isolation.

To start Storybook:

**npm run storybook**
Visit http://localhost:6006/ to view the Storybook environment.

4️⃣ ### Visual Regression with Chromatic
Chromatic ensures your UI stays consistent across changes with automated visual snapshots.

First, install Chromatic:

**npm install --save-dev chromatic**
To run Chromatic visual tests:

**npx chromatic --project-token=YOUR_PROJECT_TOKEN**
Replace YOUR_PROJECT_TOKEN with your actual Chromatic project token.

5️⃣ ### Cypress End-to-End Testing
We leverage Cypress for automated E2E tests to ensure the website works flawlessly in real-world scenarios.

Start your development server:

**npm run dev**
Run Cypress tests in interactive mode:


**npx cypress open**
Or run Cypress tests in headless mode:


**npx cypress run**
6️⃣ ### Combined Workflow with Cypress + Chromatic
We use GitHub Actions to run Cypress E2E tests alongside Chromatic visual regression in one workflow to streamline our CI/CD pipeline. Below is an overview of how our combined testing works:

Run Cypress E2E tests:

**npx cypress run**
Run Chromatic + Cypress together:

**npx chromatic --cypress --project-token=YOUR_PROJECT_TOKEN**
