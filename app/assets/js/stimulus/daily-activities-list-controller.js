import { Controller } from "@hotwired/stimulus";

/**
 * Controller for managing the activity list.
 */
export default class extends Controller {
    //static targets = ["list"];

    refreshList() {
        console.log('refreshList', event)
    }

    handleActivitySaved(event) {
        // Example: Update the list with new data
        console.log('handleActivitySaved', event)
    }
}
