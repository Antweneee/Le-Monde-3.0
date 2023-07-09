describe('[LAYOUT] Nouvel article page', () => {
	beforeEach('Go to nouvel-article page', () => {
		cy.visit('/nouvel-article');
	});

	it('Title input', () => {
		cy.get('#nouvel-article-title-input').should('be.visible');
	});

	it('Topic input', () => {
		cy.get('#nouvel-article-topic-input').should('be.visible');
	});

	it('Content textarea', () => {
		cy.get('#nouvel-article-content-textarea').should('be.visible');
	});

	it('Publish button', () => {
		cy.get('#nouvel-article-publish-btn').should('contain', 'Publier');
	});

	it('Pre-visualize button', () => {
		cy.get('#nouvel-article-pre-visualize-btn').should('contain', 'Pré-visualisez votre article');
	});

	it('Save to draft button', () => {
		cy.get('#nouvel-article-save-draft-btn').should('contain', 'Enregistrer dans les brouillons');
	});
});

describe('[LINKS] Nouvel article page', () => {
	beforeEach('Go to nouvel-article page', () => {
		cy.visit('/nouvel-article');
	});

	it('Link to publications page', () => {
		cy.get('#nouvel-article-publish-btn').click().url().should('eq', `${Cypress.config().baseUrl}/publications`);
	});

	it('Link to brouillons page', () => {
		cy.get('#nouvel-article-save-draft-btn').click().url().should('eq', `${Cypress.config().baseUrl}/brouillons`);
	});
});
