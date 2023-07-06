import pytest

from testsuite.databases import pgsql


@pytest.mark.parametrize(
    'request_body, expected_response_body, '
    'expected_response_code, expected_db_data',
    [
        pytest.param(
            {
                'customerId': 1,
                'start': {
                    'lon': 22,
                    'lat': 11,
                },
                'finish': {
                    'lon': 23,
                    'lat': 12,
                },
            },
            {
                'order_id': 1,
            },
            200,
            [(1, 1, 'new', '(11,22)', '(12,23)')],
            id='OK',
        ),
        pytest.param(
            {
                'customerId': 1,
                'start': {
                    'lon': 22,
                    'lat': 11,
                },
                'finidsh': {
                    'lon': 23,
                    'lat': 12,
                },
            },
            {'message': 'Key \'finish\' is missing but required'},
            400,
            None,
            id='error in request',
        ),
        pytest.param(
            {
                'customerId': 1,
                'start': {
                    'lon': 22,
                    'lat': 11,
                },
                'finish': {
                    'lon': 23,
                    'lat': 91,
                },
            },
            {'message': 'lat param out of bounds'},
            400,
            None,
            id='lat param out of bounds',
        ),
    ],
)
async def test_v1_order_post(service_client, request_body,
                             expected_response_body,
                             expected_response_code,
                             expected_db_data, pgsql, ):
    response = await service_client.post(
        '/v1/order',
        json=request_body,
    )
    assert response.status == expected_response_code
    assert response.json() == expected_response_body

    if expected_response_code == 200:
        cursor = pgsql['db_1'].cursor()
        cursor.execute(
            """
            select id, customer, status, start_point, end_point
            from deli_main.orders
            """,
        )
        data = list(cursor)
        assert data == expected_db_data
