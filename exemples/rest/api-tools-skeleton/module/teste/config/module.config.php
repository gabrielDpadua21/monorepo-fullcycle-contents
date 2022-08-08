<?php
return [
    'router' => [
        'routes' => [
            'teste.rest.user' => [
                'type' => 'Segment',
                'options' => [
                    'route' => '/user[/:user_id]',
                    'defaults' => [
                        'controller' => 'teste\\V1\\Rest\\User\\Controller',
                    ],
                ],
            ],
        ],
    ],
    'api-tools-versioning' => [
        'uri' => [
            0 => 'teste.rest.user',
        ],
    ],
    'api-tools-rest' => [
        'teste\\V1\\Rest\\User\\Controller' => [
            'listener' => 'teste\\V1\\Rest\\User\\UserResource',
            'route_name' => 'teste.rest.user',
            'route_identifier_name' => 'user_id',
            'collection_name' => 'user',
            'entity_http_methods' => [
                0 => 'GET',
                1 => 'PATCH',
                2 => 'PUT',
                3 => 'DELETE',
            ],
            'collection_http_methods' => [
                0 => 'GET',
                1 => 'POST',
            ],
            'collection_query_whitelist' => [],
            'page_size' => 25,
            'page_size_param' => null,
            'entity_class' => \teste\V1\Rest\User\UserEntity::class,
            'collection_class' => \teste\V1\Rest\User\UserCollection::class,
            'service_name' => 'user',
        ],
    ],
    'api-tools-content-negotiation' => [
        'controllers' => [
            'teste\\V1\\Rest\\User\\Controller' => 'HalJson',
        ],
        'accept_whitelist' => [
            'teste\\V1\\Rest\\User\\Controller' => [
                0 => 'application/vnd.teste.v1+json',
                1 => 'application/hal+json',
                2 => 'application/json',
            ],
        ],
        'content_type_whitelist' => [
            'teste\\V1\\Rest\\User\\Controller' => [
                0 => 'application/vnd.teste.v1+json',
                1 => 'application/json',
            ],
        ],
    ],
    'api-tools-hal' => [
        'metadata_map' => [
            \teste\V1\Rest\User\UserEntity::class => [
                'entity_identifier_name' => 'id',
                'route_name' => 'teste.rest.user',
                'route_identifier_name' => 'user_id',
                'hydrator' => \Laminas\Hydrator\ArraySerializable::class,
            ],
            \teste\V1\Rest\User\UserCollection::class => [
                'entity_identifier_name' => 'id',
                'route_name' => 'teste.rest.user',
                'route_identifier_name' => 'user_id',
                'is_collection' => true,
            ],
        ],
    ],
    'api-tools' => [
        'db-connected' => [
            'teste\\V1\\Rest\\User\\UserResource' => [
                'adapter_name' => 'dummy',
                'table_name' => 'user',
                'hydrator_name' => \Laminas\Hydrator\ArraySerializable::class,
                'controller_service_name' => 'teste\\V1\\Rest\\User\\Controller',
                'entity_identifier_name' => 'id',
            ],
        ],
    ],
];
