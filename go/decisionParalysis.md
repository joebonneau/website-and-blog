We've all been there... you've just spent hours customizing your terminal configuration and it's exactly how you wanted it. Sweet, you're feeling good about yourself. A couple of hours later, you read about a hot new terminal (written in Rust, of course) that claims to do so much more and is rapidly gaining popularity. You eagerly abandon all of that time you invested into your old mainstay to get a taste of the newness.

I fall into these same patterns all of the time and it can be incredibly anxiety-inducing to feel like you constantly need to evolve your development environment. I think a lot of this comes with the territory of technology, given how quickly it evolves. No one wants to feel like they've been left behind or that their skills are now outdated. I've been thinking a lot lately about how I can avoid falling into this trap.

I'm going to be focusing on enhancements to your development environment in this article. Many of the same arguments will likely apply to other areas, but I wanted to lead with that. Before you evaluate a new technology, it's important to ask yourself "what problem am I looking to solve?". Then, you might ask yourself "how much do I actually care about the problem I'm trying to solve?". Here's an example: you've just installed and configured Alacritty, a GPU-accelerated terminal emulator written in Rust. You're feeling pretty good about yourself, but then you read about a new terminal emulator called Kitty that is also GPU-accelerated and written in Rust. It's very fast and does everything you want it to. But, you read that it doesn't support ligatures. You immediately decide that you must switch to something like Wezterm because it natively supports ligatures, abandoning your Alacritty setup that you spent hours on.

Let's take a step back. First, do you even care about ligatures? Do you even know what they are? If the answer to either of those questions is "no", then it sounds like you were attracted by something new and shiny because it boasts a feature that your setup currently does not. Fundamentally, there's nothing wrong with that. One of the coolest things about software is that it's ever-evolving and innovation is exciting. But, if your current setup covers all of the bases that you currently care about, then why abandon your hard work? For me, the real-world effect of this is that I get overwhelmed about all of the options out there. I don't wantto sink a bunch of time into something to just find out there may have been a better option. However, I'd also like to spend my free time creating things, and constantly configuring my development environment will surely be a hindrance to that.

A great example of a technology that has stood the test of time is `tmux`. Tmux is a "terminal multiplexer" that allows you to maintain and switch between multiple programs in a single terminal emulator. It works great and it has for a long time. So, let's say that you've incorporated tmux into your workflow and it's been amainstay for a very long time. A new terminal multiplexer, Zellij, hits the scene and you're naturally interested to see what all the fuss is about. Do you give Zellij a try or do you put your blinders on because you're happy with tmux? The decision is ultimately yours, but it's important to consider your motivations. Is there something about tmux that has bothered you for years that you think Zellij might provide a better solution for? Or, are you just curious to see if Zellij approaches something in a way you hadn't even considered and that it could be a boon to your workflow? Either answer is totally fine - but before you throw your beloved tmux away, consider what you're looking to accomplish.