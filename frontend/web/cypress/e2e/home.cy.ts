describe('[LAYOUT] Home page', () => {
	beforeEach('Go to home page', () => {
		cy.visit('/');
	});

	it('Title', () => {
		cy.get('#app-title').should('contain', 'Le Monde 3.0');
	});

	it('Description', () => {
		cy.get('#app-description').should('contain', 'Le journal décentralisé luttant contre la censure.');
	});

	it('Connexion button', () => {
		cy.get('#home-connexion-btn').should('contain', 'Connexion');
	});

	it('Inscription button', () => {
		cy.get('#home-inscription-btn').should('contain', 'Inscription');
	});
});

describe('[LINKS] Home page', () => {
	beforeEach('Go to home page', () => {
		cy.visit('/');
	});

	it('Link to connexion page', () => {
		cy.get('#home-connexion-btn').click().url().should('eq', `${Cypress.config().baseUrl}/connexion`);
	});

	it('Link to inscription page', () => {
		cy.get('#home-inscription-btn').click().url().should('eq', `${Cypress.config().baseUrl}/inscription`);
	});
});
