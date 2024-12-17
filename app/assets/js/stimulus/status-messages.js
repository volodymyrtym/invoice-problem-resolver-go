/**
 *
 * @param {HTMLButtonElement} targetElement
 *
 * @return void
 */
export function showLoadingMessage(targetElement) {
    resetMessageClasses(targetElement); // Скидаємо всі класи
    targetElement.classList.add("loading-message");
    targetElement.style.display = "block";
    targetElement.innerHTML = `Loading<span class="dots">...</span>`;
}

/**
 *
 * @param {HTMLButtonElement} targetElement
 * @param {string} message
 *
 * @return void
 */
export function showErrorMessage(targetElement, message) {
    resetMessageClasses(targetElement); // Скидаємо всі класи
    targetElement.classList.add("error-message");
    targetElement.style.display = "block";
    targetElement.textContent = message;
}

/**
 *
 * @param {HTMLButtonElement} targetElement
 * @param {string} message
 * @return void
 */
export function showSuccessMessage(targetElement, message) {
    resetMessageClasses(targetElement); // Скидаємо всі класи
    targetElement.classList.add("success-message");
    targetElement.style.display = "block";
    targetElement.textContent = message;
}

function resetMessageClasses(targetElement) {
    targetElement.classList.remove("loading-message", "error-message", "success-message");
}

export function hideMessage(targetElement) {
    targetElement.style.display = "none";
}
