import {Controller} from "@hotwired/stimulus";
import {showErrorMessage, showLoadingMessage, showSuccessMessage,} from "./status-messages";

/**
 * @property {HTMLInputElement} emailTarget
 * @property {HTMLInputElement} passwordTarget
 * @property {HTMLElement} statusMessageTarget
 * @property {HTMLButtonElement} submitButtonTarget
 */
export default class extends Controller {
    static targets = ["email", "password", "statusMessage", "submitButton"];

    async handleSubmit(event) {
        event.preventDefault();

        const email = this.emailTarget.value;
        const password = this.passwordTarget.value;

        showLoadingMessage(this.statusMessageTarget);

        try {
            const response = await fetch("/users/login", {
                method: "PUT",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, password }),
            });

            if (!response.ok) {
                const data = await response.json();
                showErrorMessage(this.statusMessageTarget, data.error || "An error occurred");
            } else {
                showSuccessMessage(this.statusMessageTarget, "Login successful!");
                setTimeout(() => {
                    window.location.href = "/daily-activities/index";
                }, 1000);
            }
        } catch (error) {
            showErrorMessage(this.statusMessageTarget, "A network error occurred. Please try again.");
        } finally {
            this.submitButtonTarget.disabled = false;
        }
    }
}
