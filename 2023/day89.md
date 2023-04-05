# Oops: When something goes wrong - Post Mortems

As we have already discussed things will go wrong! We now know how to build applications where the impact of things breaking is less and where we know when things break. However, there will always be unknowns in the system and there will always be human error.

We need to build up a culture that fosters learning and improving from mistakes and one that understands that failure is inevitable. It’s the only way we learn. If the culture does not recognise this then we can quickly get into a situation where everyone is scared to make changes or take risks and therefore nothing improves. 

## Blameless culture

We are trying to build a “Blameless culture” where people are free to experiment, change and adapt. The key to this culture is always assuming someone had the best intentions based on the available information. This means that if somebody pressed a big button on your backend dashboard that dropped your production database they only did this because they did not have enough information to know not to do this! It means not blaming the intern for taking down production (As, why did they have the ability to make such a change?). 

Having a blameless culture leads to “Psychological Safety”. Psychological safety is a term used to describe a workplace environment where employees feel comfortable and safe to take interpersonal risks, such as sharing ideas, asking for help, or admitting mistakes, without fear of negative consequences.
In a psychologically safe workplace, employees feel that their opinions and contributions are valued, and they are free to express themselves without fear of ridicule, retribution, or ostracism. This can lead to increased trust and collaboration among team members, as well as improved innovation and problem-solving.
Psychological safety has been found to be an important factor in team effectiveness and organisational performance, and it has been linked to positive outcomes such as higher job satisfaction, lower turnover rates, and better overall team performance.
Can you see how important it is to have a blameless culture with psychological safety? Without it nothing gets done.

## Post Mortems 

Hopefully this section will explain why we have looked into a blameless culture with psychological safety. When things go wrong we need to be able to look at the objectively and ask “What went wrong? Why and What can we do to reduce the likelihood (or risk) of this happening again?”

If your team or organisation are not mature enough to sit down together and admit their mistakes, what information was missing to them at 2AM when they were trying to fix a system outage or to take the emotion of failure out of the picture then its going to be hard for everyone to share honestly. 

Can you say that your organisation would be happy to hear from a new employee on what mistake they made that had a big customer impact without blaming them? They should. If they made that mistake then it means that anyone was able to make it, at any time. Build guard rails into your system around these incidents. [This is covered well by this website covering retrospectives](https://www.funretrospectives.com/the-retrospective-prime-directive/)  as at the end of the day, a post mortem is just a retrospective. 

So, after an incident its very important to sit down and have an honest and open conversation about what went wrong and what should be done to fix those issues, and who is responsible for making those changes. This can be as informal as having a coffee with the key stakeholders or as formal as putting a calendar invite in and getting everyone together from across the business. 

