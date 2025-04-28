using System.Text.Json;
using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Mvc;

namespace Subscriptions.Api.Controllers;

[ApiController]
[Route("api/subscriptions")]
public class SubscriptionValidationController(
    ILogger<SubscriptionValidationController> logger,
    HttpClient httpClient)
    : ControllerBase
{
    [HttpGet(template: "/active")]
    public async Task<ActionResult<bool>> Get(long groupId, long userId)
    {
        logger.LogInformation("Verifying subscription {subscriptionId} for group {groupId}", userId, groupId);
        
        try
        {
            var response = await httpClient.GetAsync(
                $"https://api.vk.com/method/groups.isMember?group_id={groupId}&user_id={userId}&v=5.199");
            response.EnsureSuccessStatusCode();

            var responseBody = JsonSerializer.Deserialize<IsMemberResponse>(await response.Content.ReadAsStringAsync());
            return responseBody!.IsSubscribed;
        }
        catch (Exception ex)
        {
            logger.LogError(
                ex,
                "Unknown error occured while fetching subscription info for group {groupId} and user {userId}",
                groupId,
                userId);
            return StatusCode(500);
        }
    }
    
    private class IsMemberResponse
    {
        [JsonPropertyName("response")]
        public int Response { private get; init; }
        
        public bool IsSubscribed => Response == 1;
    }
}