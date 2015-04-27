function supportsLocalStorage() {

	
    return ('localStorage' in window) && window['localStorage'] !== null;
}