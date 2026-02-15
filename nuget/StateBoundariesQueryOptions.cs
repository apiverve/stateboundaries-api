using System;
using System.Collections.Generic;
using System.Text;
using Newtonsoft.Json;

namespace APIVerve.API.StateBoundaries
{
    /// <summary>
    /// Query options for the State Boundaries API
    /// </summary>
    public class StateBoundariesQueryOptions
    {
        /// <summary>
        /// The 2 Letter State Code (e.g., CA)
        /// </summary>
        [JsonProperty("state")]
        public string State { get; set; }
    }
}
