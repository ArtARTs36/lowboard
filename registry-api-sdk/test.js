import { SDK } from './dist/index.js';

const sdk = new SDK({
    baseUrl: 'http://localhost:8000',
});

sdk.client.getDefinition()
    .then((p) => {
        console.log(p.pages[0].name)
    })
