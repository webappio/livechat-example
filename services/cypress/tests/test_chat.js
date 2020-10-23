describe('The Chat Page', () => {
    beforeEach(() => {
        const username = "testuser"
        const password = "p@ssword"

        const formData = new FormData();
        formData.append('name', 'myuser');
        formData.append('password', 'p@assword');
        cy.request({
            method: 'POST',
            url: '/api/login',
            form: true,
            body: {'name': 'myuser', 'password': 'p@assword'},
        });

        // our auth cookie should be present
        cy.getCookie('default').should('exist');
    })

    it('can send a chat message', function () {
        cy.visit('/');

        let num = Math.random();
        cy.get('.message-area').type(`some message #${num}{enter}`);

        cy.get('.message-wrapper:last-child .message-text').should('have.text', `some message #${num}`)
    })
})