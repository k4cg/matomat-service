# coding: utf-8

"""
    MaaS

    Draft for MaaS (Matomat as a Service)  # noqa: E501

    OpenAPI spec version: 0.2.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from __future__ import absolute_import

import re  # noqa: F401

# python 2 and python 3 compatibility library
import six

from swagger_client.api_client import ApiClient


class ItemsApi(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    Ref: https://github.com/swagger-api/swagger-codegen
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client

    def items_get(self, **kwargs):  # noqa: E501
        """List all available items  # noqa: E501

        Returns a map of item objects, with the item ID as key and the object as value  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_get(async=True)
        >>> result = thread.get()

        :param async bool
        :return: list[Item]
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async'):
            return self.items_get_with_http_info(**kwargs)  # noqa: E501
        else:
            (data) = self.items_get_with_http_info(**kwargs)  # noqa: E501
            return data

    def items_get_with_http_info(self, **kwargs):  # noqa: E501
        """List all available items  # noqa: E501

        Returns a map of item objects, with the item ID as key and the object as value  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_get_with_http_info(async=True)
        >>> result = thread.get()

        :param async bool
        :return: list[Item]
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = []  # noqa: E501
        all_params.append('async')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method items_get" % key
                )
            params[key] = val
        del params['kwargs']

        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/items', 'GET',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='list[Item]',  # noqa: E501
            auth_settings=auth_settings,
            async=params.get('async'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def items_item_id_consume_put(self, item_id, **kwargs):  # noqa: E501
        """Consumes a Item  # noqa: E501

        Consumes a Item and subtracts the cost of the Item from the credit of the user. If not enough credit exists the operation will be rejected  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_item_id_consume_put(item_id, async=True)
        >>> result = thread.get()

        :param async bool
        :param int item_id: The ID of the Item that needs to be consumed (required)
        :return: ItemStats
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async'):
            return self.items_item_id_consume_put_with_http_info(item_id, **kwargs)  # noqa: E501
        else:
            (data) = self.items_item_id_consume_put_with_http_info(item_id, **kwargs)  # noqa: E501
            return data

    def items_item_id_consume_put_with_http_info(self, item_id, **kwargs):  # noqa: E501
        """Consumes a Item  # noqa: E501

        Consumes a Item and subtracts the cost of the Item from the credit of the user. If not enough credit exists the operation will be rejected  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_item_id_consume_put_with_http_info(item_id, async=True)
        >>> result = thread.get()

        :param async bool
        :param int item_id: The ID of the Item that needs to be consumed (required)
        :return: ItemStats
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['item_id']  # noqa: E501
        all_params.append('async')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method items_item_id_consume_put" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'item_id' is set
        if ('item_id' not in params or
                params['item_id'] is None):
            raise ValueError("Missing the required parameter `item_id` when calling `items_item_id_consume_put`")  # noqa: E501

        collection_formats = {}

        path_params = {}
        if 'item_id' in params:
            path_params['itemId'] = params['item_id']  # noqa: E501

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/items/{itemId}/consume', 'PUT',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='ItemStats',  # noqa: E501
            auth_settings=auth_settings,
            async=params.get('async'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def items_item_id_delete(self, item_id, **kwargs):  # noqa: E501
        """Delete Item  # noqa: E501

        Delete the Item. This can only be done by admins. (Should only mark a Item as deleted to not loose reference for stats)  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_item_id_delete(item_id, async=True)
        >>> result = thread.get()

        :param async bool
        :param int item_id: The ID of the Item to perform an operation on (required)
        :return: Item
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async'):
            return self.items_item_id_delete_with_http_info(item_id, **kwargs)  # noqa: E501
        else:
            (data) = self.items_item_id_delete_with_http_info(item_id, **kwargs)  # noqa: E501
            return data

    def items_item_id_delete_with_http_info(self, item_id, **kwargs):  # noqa: E501
        """Delete Item  # noqa: E501

        Delete the Item. This can only be done by admins. (Should only mark a Item as deleted to not loose reference for stats)  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_item_id_delete_with_http_info(item_id, async=True)
        >>> result = thread.get()

        :param async bool
        :param int item_id: The ID of the Item to perform an operation on (required)
        :return: Item
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['item_id']  # noqa: E501
        all_params.append('async')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method items_item_id_delete" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'item_id' is set
        if ('item_id' not in params or
                params['item_id'] is None):
            raise ValueError("Missing the required parameter `item_id` when calling `items_item_id_delete`")  # noqa: E501

        collection_formats = {}

        path_params = {}
        if 'item_id' in params:
            path_params['itemId'] = params['item_id']  # noqa: E501

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/items/{itemId}', 'DELETE',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='Item',  # noqa: E501
            auth_settings=auth_settings,
            async=params.get('async'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def items_item_id_get(self, item_id, **kwargs):  # noqa: E501
        """Get a certain Item  # noqa: E501

        Get a certain Item  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_item_id_get(item_id, async=True)
        >>> result = thread.get()

        :param async bool
        :param int item_id: The ID of the Item to perform an operation on (required)
        :return: Item
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async'):
            return self.items_item_id_get_with_http_info(item_id, **kwargs)  # noqa: E501
        else:
            (data) = self.items_item_id_get_with_http_info(item_id, **kwargs)  # noqa: E501
            return data

    def items_item_id_get_with_http_info(self, item_id, **kwargs):  # noqa: E501
        """Get a certain Item  # noqa: E501

        Get a certain Item  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_item_id_get_with_http_info(item_id, async=True)
        >>> result = thread.get()

        :param async bool
        :param int item_id: The ID of the Item to perform an operation on (required)
        :return: Item
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['item_id']  # noqa: E501
        all_params.append('async')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method items_item_id_get" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'item_id' is set
        if ('item_id' not in params or
                params['item_id'] is None):
            raise ValueError("Missing the required parameter `item_id` when calling `items_item_id_get`")  # noqa: E501

        collection_formats = {}

        path_params = {}
        if 'item_id' in params:
            path_params['itemId'] = params['item_id']  # noqa: E501

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/items/{itemId}', 'GET',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='Item',  # noqa: E501
            auth_settings=auth_settings,
            async=params.get('async'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def items_item_id_put(self, item_id, name, cost, **kwargs):  # noqa: E501
        """Update Item  # noqa: E501

        Update Item. This can only be done by admins  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_item_id_put(item_id, name, cost, async=True)
        >>> result = thread.get()

        :param async bool
        :param int item_id: The ID of the Item to perform an operation on (required)
        :param str name: (required)
        :param int cost: (required)
        :return: Item
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async'):
            return self.items_item_id_put_with_http_info(item_id, name, cost, **kwargs)  # noqa: E501
        else:
            (data) = self.items_item_id_put_with_http_info(item_id, name, cost, **kwargs)  # noqa: E501
            return data

    def items_item_id_put_with_http_info(self, item_id, name, cost, **kwargs):  # noqa: E501
        """Update Item  # noqa: E501

        Update Item. This can only be done by admins  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_item_id_put_with_http_info(item_id, name, cost, async=True)
        >>> result = thread.get()

        :param async bool
        :param int item_id: The ID of the Item to perform an operation on (required)
        :param str name: (required)
        :param int cost: (required)
        :return: Item
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['item_id', 'name', 'cost']  # noqa: E501
        all_params.append('async')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method items_item_id_put" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'item_id' is set
        if ('item_id' not in params or
                params['item_id'] is None):
            raise ValueError("Missing the required parameter `item_id` when calling `items_item_id_put`")  # noqa: E501
        # verify the required parameter 'name' is set
        if ('name' not in params or
                params['name'] is None):
            raise ValueError("Missing the required parameter `name` when calling `items_item_id_put`")  # noqa: E501
        # verify the required parameter 'cost' is set
        if ('cost' not in params or
                params['cost'] is None):
            raise ValueError("Missing the required parameter `cost` when calling `items_item_id_put`")  # noqa: E501

        if 'cost' in params and params['cost'] < 0:  # noqa: E501
            raise ValueError("Invalid value for parameter `cost` when calling `items_item_id_put`, must be a value greater than or equal to `0`")  # noqa: E501
        collection_formats = {}

        path_params = {}
        if 'item_id' in params:
            path_params['itemId'] = params['item_id']  # noqa: E501

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}
        if 'name' in params:
            form_params.append(('name', params['name']))  # noqa: E501
        if 'cost' in params:
            form_params.append(('cost', params['cost']))  # noqa: E501

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # HTTP header `Content-Type`
        header_params['Content-Type'] = self.api_client.select_header_content_type(  # noqa: E501
            ['application/x-www-form-urlencoded'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/items/{itemId}', 'PUT',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='Item',  # noqa: E501
            auth_settings=auth_settings,
            async=params.get('async'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def items_item_id_stats_get(self, item_id, **kwargs):  # noqa: E501
        """Get consumption stats  # noqa: E501

        Get the matomat stats for a certain Item  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_item_id_stats_get(item_id, async=True)
        >>> result = thread.get()

        :param async bool
        :param int item_id: The ID of the Item for which to fetch the stats (required)
        :return: ItemStats
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async'):
            return self.items_item_id_stats_get_with_http_info(item_id, **kwargs)  # noqa: E501
        else:
            (data) = self.items_item_id_stats_get_with_http_info(item_id, **kwargs)  # noqa: E501
            return data

    def items_item_id_stats_get_with_http_info(self, item_id, **kwargs):  # noqa: E501
        """Get consumption stats  # noqa: E501

        Get the matomat stats for a certain Item  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_item_id_stats_get_with_http_info(item_id, async=True)
        >>> result = thread.get()

        :param async bool
        :param int item_id: The ID of the Item for which to fetch the stats (required)
        :return: ItemStats
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['item_id']  # noqa: E501
        all_params.append('async')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method items_item_id_stats_get" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'item_id' is set
        if ('item_id' not in params or
                params['item_id'] is None):
            raise ValueError("Missing the required parameter `item_id` when calling `items_item_id_stats_get`")  # noqa: E501

        collection_formats = {}

        path_params = {}
        if 'item_id' in params:
            path_params['itemId'] = params['item_id']  # noqa: E501

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/items/{itemId}/stats', 'GET',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='ItemStats',  # noqa: E501
            auth_settings=auth_settings,
            async=params.get('async'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def items_post(self, name, cost, **kwargs):  # noqa: E501
        """Add a new item  # noqa: E501

        Adds a new item to matomat. This can only be done by admins  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_post(name, cost, async=True)
        >>> result = thread.get()

        :param async bool
        :param str name: (required)
        :param int cost: (required)
        :return: Item
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async'):
            return self.items_post_with_http_info(name, cost, **kwargs)  # noqa: E501
        else:
            (data) = self.items_post_with_http_info(name, cost, **kwargs)  # noqa: E501
            return data

    def items_post_with_http_info(self, name, cost, **kwargs):  # noqa: E501
        """Add a new item  # noqa: E501

        Adds a new item to matomat. This can only be done by admins  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_post_with_http_info(name, cost, async=True)
        >>> result = thread.get()

        :param async bool
        :param str name: (required)
        :param int cost: (required)
        :return: Item
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['name', 'cost']  # noqa: E501
        all_params.append('async')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method items_post" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'name' is set
        if ('name' not in params or
                params['name'] is None):
            raise ValueError("Missing the required parameter `name` when calling `items_post`")  # noqa: E501
        # verify the required parameter 'cost' is set
        if ('cost' not in params or
                params['cost'] is None):
            raise ValueError("Missing the required parameter `cost` when calling `items_post`")  # noqa: E501

        if 'cost' in params and params['cost'] < 0:  # noqa: E501
            raise ValueError("Invalid value for parameter `cost` when calling `items_post`, must be a value greater than or equal to `0`")  # noqa: E501
        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}
        if 'name' in params:
            form_params.append(('name', params['name']))  # noqa: E501
        if 'cost' in params:
            form_params.append(('cost', params['cost']))  # noqa: E501

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # HTTP header `Content-Type`
        header_params['Content-Type'] = self.api_client.select_header_content_type(  # noqa: E501
            ['application/x-www-form-urlencoded'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/items', 'POST',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='Item',  # noqa: E501
            auth_settings=auth_settings,
            async=params.get('async'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def items_stats_get(self, **kwargs):  # noqa: E501
        """Get consumption stats of all items  # noqa: E501

        Get the matomat stats for all items  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_stats_get(async=True)
        >>> result = thread.get()

        :param async bool
        :return: list[ItemStats]
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async'):
            return self.items_stats_get_with_http_info(**kwargs)  # noqa: E501
        else:
            (data) = self.items_stats_get_with_http_info(**kwargs)  # noqa: E501
            return data

    def items_stats_get_with_http_info(self, **kwargs):  # noqa: E501
        """Get consumption stats of all items  # noqa: E501

        Get the matomat stats for all items  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async=True
        >>> thread = api.items_stats_get_with_http_info(async=True)
        >>> result = thread.get()

        :param async bool
        :return: list[ItemStats]
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = []  # noqa: E501
        all_params.append('async')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method items_stats_get" % key
                )
            params[key] = val
        del params['kwargs']

        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/items/stats', 'GET',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='list[ItemStats]',  # noqa: E501
            auth_settings=auth_settings,
            async=params.get('async'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)
