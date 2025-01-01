import {Controller} from "@hotwired/stimulus";
import {showErrorMessage, showLoadingMessage, showSuccessMessage,} from "./status-messages";


/**
 * Controller for managing the visibility of the activity form.
 *
 * @extends Controller

 * Stimulus Targets
 * @property {HTMLElement} statusMessageTarget
 */
export default class extends Controller {
    static targets = ["statusMessage"];

    resetForm(form) {
        form.reset(); // Resets all input, select, and textarea fields to their default values.
    }

    /**
     * Submits the form data via an HTTP request and dispatches a custom event.
     * @param {Event} event - The event triggered by clicking the Save button.
     */
    async handleSubmit(event) {
        showLoadingMessage(this.statusMessageTarget);
        event.preventDefault();

        const url = "/api/daily-activities";
        const form = event.target;
        const formData = new FormData(form);
        const jsonData = Object.fromEntries(formData.entries());

        try {
            const response = await fetch(url, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(jsonData), // Send data as JSON
            });

            if (response.ok) {
                window.location.reload();
            }

            showErrorMessage(this.statusMessageTarget, response.message || "An error occurred");
            //this.dispatch("activitySaved", {detail: {response: await response.json()}})
            //showSuccessMessage(this.statusMessageTarget, "Activity added!");
            //this.resetForm(form);
        } catch (error) {
            showErrorMessage(this.statusMessageTarget, error.message || "Critical error occurred");
        }
    }
}
