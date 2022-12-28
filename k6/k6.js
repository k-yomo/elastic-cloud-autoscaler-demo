import http from 'k6/http';
import {check, fail} from 'k6';
import { SharedArray } from 'k6/data';
import papaparse from 'https://jslib.k6.io/papaparse/5.1.1/index.js';

const csvData = new SharedArray('queries', function () {
    return papaparse.parse(open('./queries.csv'), { header: true }).data;
});

const searchQueryPayload = (query) => {
    return JSON.stringify({
        "query": {
            "wildcard": {
                "name": {
                    "value": `*${query}*`
                }
            }
        },
        "sort": [
            {
                "weight": {
                    "order": "desc"
                }
            }
        ],
        "_source": [],
        "size": 100
    })
}

export default function () {
    const randomQuery = csvData[Math.floor(Math.random() * csvData.length)]['query'];
    const payload = searchQueryPayload(randomQuery)

    const params = {
        headers: {
            "Content-Type": "application/json",
            "kbn-xsrf": "reporting",
            "Authorization": ""
        }
    }
    const res = http.post(
        ``,
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
