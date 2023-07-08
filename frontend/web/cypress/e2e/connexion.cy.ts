describe('[LAYOUT] Connexion page', () => {
	beforeEach('Go to connexion page', () => {
		cy.visit('/connexion');
	});

	it('Title', () => {
		cy.get('#app-title').should('contain', 'Le Monde 3.0');
	});

	it('Description', () => {
		cy.get('#app-description').should('contain', 'Le journal décentralisé luttant contre la censure.');
	});

	it('Email / username input', () => {
		cy.get('#connexion-email-input').should('be.visible');
	});

	it('Password input', () => {
		cy.get('#connexion-pwd-input').should('be.visible');
	});

	it('Connexion button', () => {
		cy.get('#connexion-connexion-btn').should('contain', 'Connexion');
	});

	it('Inscription button', () => {
		cy.get('#connexion-inscription-btn').should('contain', 'Inscription');
	});
});

describe('[LINKS] Connexion page', () => {
	beforeEach('Go to connexion page', () => {
		cy.visit('/connexion');
	});

	it('Link to favoris page', () => {
		cy.get('#connexion-connexion-btn').click().url().should('eq', `${Cypress.config().baseUrl}/favoris`);
	});

	it('Link to inscription page', () => {
		cy.get('#connexion-inscription-btn').click().url().should('eq', `${Cypress.config().baseUrl}/inscription`);
	});
});
