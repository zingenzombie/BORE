import { DescriptionBoxComponent } from "./description-box.component";

describe('DescriptionBoxComponent', () => {
    it('mounts', () => {
        cy.mount(DescriptionBoxComponent)
    })

    it('button should exist', () => {
        cy.mount(DescriptionBoxComponent)
        cy.get('button').should('exist')
    })

    it('button should display About', () => {
        cy.mount(DescriptionBoxComponent)
        cy.get('button').should('have.text', 'About')
        cy.get('button').should('exist')
    }) 

    it('button should toggle the description', () => {
        cy.mount(DescriptionBoxComponent)
        cy.get('button').click();
        cy.get('div').should('be.visible')
    })
})