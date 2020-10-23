it('sets auth cookie when logging in via form submission', function () {
    const username = "testuser"
    const password = "p@ssword"

    cy.visit('/login')

    cy.get('input[name=name]').type(username)

    // {enter} causes the form to submit
    cy.get('input[name=password]').type(`${password}{enter}`)

    cy.url().should('eq', 'http://traefik/')

    // our auth cookie should be present
    cy.getCookie('default').should('exist')
})