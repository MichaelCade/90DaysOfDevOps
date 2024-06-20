# Day 47 - Azure logic app, low / no code
[![Watch the video](thumbnails/day47.png)](https://www.youtube.com/watch?v=pEB4Kp6JHfI)

 It seems like you have successfully created an end-to-end workflow using Azure Logic Apps that processes a grocery receipt image, identifies food items, fetches recipes for those foods, and sends an email with the list of recipes.

To continue with the next step, follow these instructions:

1. Save your workflow in your GitHub repository (if you haven't already) so you can access it later.
2. To run the workflow, you need to authenticate each connector as mentioned during the explanation:
   - Azure Blob Storage: You will need to provide authentication for the storage account where the receipt image is stored.
   - Computer Vision API (OCR): Provide authentication for your Computer Vision resource.
   - Outlook API: Authenticate with your Outlook account to send emails.
3. To test the workflow, upload a new grocery receipt image in the specified storage account.
4. Wait for an email with the list of potential recipes based on the items detected in the receipt.
5. Review and make changes as needed to improve the workflow or add more features (such as adding JavaScripts, Python functions, etc.).
6. Share your experiences, improvements, feedback, and new ideas using Azure Logic Apps in the comments section.
7. Enjoy learning and exploring the possibilities of this powerful tool!
In this session, we explored creating a workflow using Azure Logic Apps with minimal code knowledge. The goal was to automate a process that takes a receipt as input, extracts relevant information, and sends an email with potential recipes based on the food items purchased.

The workflow consisted of several steps:

1. Blob Trigger: A blob trigger was set up to capture new receipts uploaded to a storage account.
2. JSON Output: The receipt content was passed through OCR (Optical Character Recognition) and computer vision, which converted the text into a JSON format.
3. Schema Classification: The JSON output was then classified using a schema, allowing us to extract specific properties or objects within the JSON.
4. Filtering and Looping: An array of food-related texts was created by filtering the original JSON output against a food word list. A loop was used to iterate through each recipe, extracting its name, URL, and image (or thumbnail).
5. Email Body: The email body was constructed using variables for the food labels and URLs, listing out potential recipes for the user.

The final step was sending the email with the recipe list using the Outlook connector.

Key takeaways from this session include:

* Azure Logic Apps can be used to simplify workflows without requiring extensive coding knowledge.
* The platform provides a range of connectors and actions that can be combined to achieve specific business outcomes.
* Creativity and experimentation are encouraged, as users can add their own custom code snippets or integrate with other services.

The GitHub repository accompanying this session provides the complete code view of the workflow, allowing users to copy and modify it for their own purposes.
