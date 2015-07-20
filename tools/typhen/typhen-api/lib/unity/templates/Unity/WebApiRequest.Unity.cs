﻿using UnityEngine;
using System;
using System.Collections;

namespace TyphenApi
{
    public partial class WebApiRequest<ResponseT, ErrorT>
        where ResponseT : TypeBase, new()
        where ErrorT : TypeBase, new()
    {
        public Coroutine Send(Func<IEnumerator, Coroutine> startCorotineFunc,
            Action<WebApiResponse<ResponseT>> onSuccess,
            Action<WebApiError<ErrorT>> onError = null)
        {
            return startCorotineFunc(SendAsync(startCorotineFunc, onSuccess, onError));
        }

        IEnumerator SendAsync(Func<IEnumerator, Coroutine> startCorotineFunc,
            Action<WebApiResponse<ResponseT>> onSuccess,
            Action<WebApiError<ErrorT>> onError = null)
        {
            yield return startCorotineFunc(SendAsync());

            if (Error == null)
            {
                if (onSuccess != null)
                {
                    onSuccess(Response);
                }
            }
            else
            {
                if (onError != null)
                {
                    onError(Error);
                }
            }
        }
    }
}
