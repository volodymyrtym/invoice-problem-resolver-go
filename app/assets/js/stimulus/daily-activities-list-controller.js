import { Controller } from "@hotwired/stimulus";
import {loadingButtonDone, loadingButtonStart} from "./status-messages";

/**
 * Controller for managing the activity list.
 *
 * Stimulus Targets
 * @property {HTMLElement} statusMessageTarget
 *
 */
export default class extends Controller {
    /**
     * Deletes a record via an HTTP request and dispatches a custom event.
     * @param {Event} event - The event triggered by clicking the Delete button.
     */
    async handleDelete(event) {

        event.preventDefault();
        loadingButtonStart(event.target)
        const activityId = event.target.dataset.id;
        const url = `/api/daily-activities/${activityId}`;

        console.log('handleDelete', activityId, url)
        try {
            const response = await fetch(url, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                },
            });

            if (response.ok) {
                window.location.reload();

                return
            }

            alert(response.message || "An error occurred")
            loadingButtonDone(event.target)
        } catch (error) {
            alert(error.message || "An error occurred")
            loadingButtonDone(event.target)
        }
    }

    /**
     * Handles the "Previous" button click.
     * Decreases the current page parameter and refreshes the list.
     */
    goToPreviousPage() {
        const currentPage = this.getCurrentPage();
        if (currentPage > 1) {
            this.updatePageParameter(currentPage - 1);
        }
    }

    /**
     * Handles the "Next" button click.
     * Increases the current page parameter and refreshes the list.
     */
    goToNextPage() {
        const currentPage = this.getCurrentPage();
        this.updatePageParameter(currentPage + 1);
    }

    /**
     * Gets the current page from the URL.
     * @returns {number} The current page number.
     */
    getCurrentPage() {
        const urlParams = new URLSearchParams(window.location.search);

        return parseInt(urlParams.get("page") || "1", 10);
    }


    /**
     * Updates the "page" parameter in the URL with a new value and reloads the page.
     * @param {number} newPage - The new page number to set in the URL.
     */
    updatePageParameter(newPage) {
        const urlParams = new URLSearchParams(window.location.search);
        urlParams.set("page", newPage.toString());
        window.location.search = urlParams.toString(); // Updates the URL and reloads the page
    }
    
    refreshList() {
        console.log('refreshList', event)
    }

    handleActivitySaved(event) {
        // Example: Update the list with new data
        console.log('handleActivitySaved', event)
    }
}
