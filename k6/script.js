import http from 'k6/http';
import {check, fail} from 'k6';
import { SharedArray } from 'k6/data';
import papaparse from 'https://jslib.k6.io/papaparse/5.1.1/index.js';
import encoding from 'k6/encoding';

export default function () {
    const randomQuery = queriesCSV[Math.floor(Math.random() * queriesCSV.length)]['query'];
    const payload = searchQueryPayload(randomQuery)
    const base64EncodedAuthKey = encoding.b64encode(`${__ENV.ELASTICSEARCH_USERNAME}:${__ENV.ELASTICSEARCH_PASSWORD}`)
    const params = {
        headers: {
            "Content-Type": "application/json",
            "kbn-xsrf": "reporting",
            "Authorization": `Basic ${base64EncodedAuthKey}`
        }
    }
    const res = http.post(
        `${__ENV.ELASTICSEARCH_URL}/products/_search`,
        payload,
        params
    )
    if (
        !check(res, {
            'status: 200': (res) => res.status === 200,
        })
    ) {
        fail(`status: ${res.status}`);
    }
}

const queriesCSV = new SharedArray('queries', function () {
    return papaparse.parse(open('./queries.csv'), { header: true }).data;
});

const searchQueryPayload = (query) => {
    // using wildcard query to increase CPU util
    return JSON.stringify({
        "query": {
            "wildcard": {
                "_all": {
                    "value": `*${query}*`
                }
            }
        },
        "_source": [],
        "size": 100
    })
}
