describe('the links on the website', () => {
    const visitLinkClosure = (startingUrls, done) => {
        let state = {
            visited: {},
            seen: {},
            queue: startingUrls
        }
        const v = () => {
            if(!state.queue.length) {
                done();
            } else {
                let newLoc = state.queue.pop();
                state.visited[newLoc] = true;
                return cy.visit(newLoc).location().then((loc) => {
                    let baseurl = loc.href.substring(0, loc.href.length - loc.pathname.length);
                    let basepath;
                    if (loc.pathname.indexOf("/") > -1) {
                        basepath = loc.pathname.substring(0, loc.pathname.lastIndexOf("/") + 1);
                    } else {
                        basepath = "/";
                    }
                    cy.get('body').then(($body) => {
                        // check for links on page
                        // (without $body.find it errors here - cy.get expects the elements to exist)
                        if ($body.find('a:not(.no-href-check)').length) {
                            cy.get('a:not(.no-href-check)').each(link => {
                                let href = link.attr("href");
                                expect(href+','+link.attr("class")).to.not.contain("no-href-check")
                                if (!href.startsWith("/")) {
                                    href = basepath + href;
                                }
                                state.seen[href] = true;
                                state.queue.push(href);
                            });
                        }
                    })
                }).then(v);
            }
        }
        return v();
    }
    it('should not be broken when logged out', (done) => {
        return visitLinkClosure(['/'], done);
    })
})