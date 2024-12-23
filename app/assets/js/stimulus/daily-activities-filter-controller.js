import { Controller } from "@hotwired/stimulus";
import { startOfMonth, endOfMonth, subMonths, format } from "date-fns";

/**
 * Stimulus Controller responsible for handling date range updates and form submission
 * for the previous month and the current month.
 *
 * @property {HTMLElement} formTarget
 * @property {HTMLElement} startDateTarget
 * @property {HTMLElement} endDateTarget
 *
 * @extends Controller
 */
export default class extends Controller {
    static targets = ["form", "startDate", "endDate"];

    connect() {
        this.setDatesFromUrl();
    }

    setDatesFromUrl() {
        const urlParams = new URLSearchParams(window.location.search);

        const startDate = urlParams.get("start_date");
        const endDate = urlParams.get("end_date");

        if (startDate) {
            this.startDateTarget.value = startDate;
        }

        if (endDate) {
            this.endDateTarget.value = endDate;
        }
    }

    previousMonth() {
        const { startOfPreviousMonth, endOfPreviousMonth } = this.calculatePreviousMonth();

        this.startDateTarget.value = startOfPreviousMonth;
        this.endDateTarget.value = endOfPreviousMonth;

        this.formTarget.submit();
    }
    
    currentMonth() {
        const { startOfCurrentMonth, today } = this.calculateCurrentMonth();

        this.startDateTarget.value = startOfCurrentMonth;
        this.endDateTarget.value = today;

        this.formTarget.submit();
    }

    calculatePreviousMonth() {
        const now = new Date();

        // Розрахунок початку і кінця попереднього місяця
        const startOfPreviousMonth = format(startOfMonth(subMonths(now, 1)), "yyyy-MM-dd");
        const endOfPreviousMonth = format(endOfMonth(subMonths(now, 1)), "yyyy-MM-dd");

        return { startOfPreviousMonth, endOfPreviousMonth };
    }

    calculateCurrentMonth() {
        const now = new Date();

        // Початок поточного місяця і сьогоднішній день
        const startOfCurrentMonth = format(startOfMonth(now), "yyyy-MM-dd");
        const today = format(now, "yyyy-MM-dd");

        return { startOfCurrentMonth, today };
    }
}